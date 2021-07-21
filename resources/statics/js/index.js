const app = Vue.createApp({
    setup() {
        Vue.onMounted(() => {
            console.log('launch on web browser.')
        })
    }
}).use(ElementPlus).mount('#app')
