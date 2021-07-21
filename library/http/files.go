package http

import (
	"bytes"
	"errors"
	"io"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
)

type FileMap map[string][]string

type Files struct {
	files           FileMap
	count           int
	writer          *multipart.Writer
	ExtraParams     HeaderMap
	fileBuffer      *bytes.Buffer
	fileReadBuffer  *io.PipeReader
	fileWriteBuffer *io.PipeWriter
}

func (f *Files) SetFiles(fs FileMap) *Files {
	f.files = fs
	return f
}

func (f *Files) GetFiles() FileMap {
	return f.files
}

func (f *Files) GetPipeReader() *io.PipeReader {
	return f.fileReadBuffer
}

func (f *Files) GetPipeWriter() *io.PipeWriter {
	return f.fileWriteBuffer
}

func (f *Files) GetPipe() (*io.PipeReader, *io.PipeWriter) {
	return f.fileReadBuffer, f.fileWriteBuffer
}

func (f *Files) GetWriter() *multipart.Writer {
	return f.writer
}

func (f *Files) CountFile() int {
	return f.count
}

func (f *Files) HasFile() bool {
	return f.CountFile() > 0
}

func (Files) FileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

func (f *Files) loadFileToPipeBuffer(name, path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	part, err := f.writer.CreateFormFile(name, filepath.Base(path))
	if err != nil {
		return err
	}
	if _, err = io.Copy(part, file); err != nil {
		return err
	}
	return nil
}

func (f *Files) partFileToPipe() *Files {
	defer f.fileWriteBuffer.Close()
	// added extra params
	if f.ExtraParams != nil {
		for key, val := range f.ExtraParams {
			if err := f.writer.WriteField(key, val); err != nil {
				log.Println(f.fileWriteBuffer.CloseWithError(err))
			}
		}
	}
	// added file to request body
	if len(f.files) > 0 {
		for name, files := range f.files {
			for _, path := range files {
				if err := f.loadFileToPipeBuffer(name, path); err != nil {
					log.Println(err)
				}
			}
		}
	}
	return f
}

// PipeFile send big file
func (f *Files) PipeFile() func() *Files {
	// init writer body
	return f.initPipeWriterBuffer().partFileToPipe
}

// Encode if small file no problem
func (f *Files) Encode() *bytes.Buffer {
	// init writer body
	f.initWriterBuffer()
	// added extra params
	if f.ExtraParams != nil {
		for key, val := range f.ExtraParams {
			_ = f.writer.WriteField(key, val)
		}
	}
	// added file to request body
	if len(f.files) > 0 {
		for name, files := range f.files {
			for _, path := range files {
				file, err := os.Open(path)
				part, err := f.writer.CreateFormFile(name, filepath.Base(path))
				if err != nil {
					log.Fatalln(err)
				}
				if _, err = io.Copy(part, file); err != nil {
					log.Fatalln(err)
				}
			}
		}
	}
	_ = f.writer.Close()
	return f.fileBuffer
}

func (f *Files) initPipeWriterBuffer() *Files {
	if f.fileReadBuffer == nil {
		f.fileReadBuffer, f.fileWriteBuffer = io.Pipe()
	}
	if f.writer == nil {
		f.writer = multipart.NewWriter(f.fileWriteBuffer)
	}
	return f
}

func (f *Files) initWriterBuffer() {
	if f.fileBuffer == nil {
		f.fileBuffer = new(bytes.Buffer)
	}
	if f.writer == nil {
		f.writer = multipart.NewWriter(f.fileBuffer)
	}
}

func (f *Files) AddParam(key, val string) *Files {
	if f.ExtraParams == nil {
		f.ExtraParams = make(HeaderMap, 10)
	}
	f.ExtraParams[key] = val
	return f
}

func (f *Files) AddParams(params HeaderMap) *Files {
	if f.ExtraParams == nil {
		f.ExtraParams = make(HeaderMap, len(params))
	}
	for key, val := range params {
		f.ExtraParams[key] = val
	}
	return f
}

func (f *Files) AddFile(key, path string) error {
	if !f.FileExists(path) {
		return errors.New(`file not exists`)
	}
	f.count++
	if f.files == nil {
		f.files = make(map[string][]string, 10)
	}
	f.files[key] = append(f.files[key], path)
	return nil
}

func (f *Files) RemoveFile(name string) *Files {
	if _, ok := f.files[name]; ok {
		delete(f.files, name)
	}
	return f
}
