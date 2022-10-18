<template>
<div id="app" v-title data-title="redisMonitor" style="background:#14121f">
<div>
<Index></Index>
<Footer></Footer>
</div>
</div>
</template>

<script>
import Footer from './components/Footer';
import Index from './Index';

export default{
    name:'App',
    components:{
        Footer,
        Index
    },
    data(){
        return {
            content:null,
        };
    },
    methods:{
        getContext(){
            Vue.axios.get("/data").then((response)=>(this.content = response.data));
        }
    }
}
</script>


<style>

</style>