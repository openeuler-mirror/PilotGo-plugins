<template>
<div class="container" style="height:500px">
  <div class="box" style="flex-direction: row;">
                <div class="box-item" style="order: 1;flex-grow: 1;height: 500px;margin-left:10px">
                <h4 align="center">集群信息</h4>
                <div>
                <p>是否启动:</p>
                 </div>
                </div>

                 <div class="box-item" style="order: 2;flex-grow: 1;height: 500px;margin-right:10px">
                <h4 align="center">执行命令</h4>
                <div>
                <ve-bar-chart :data="chartData" :setting="chartSettings" :height="400" :textStyle="textStyle">
                </ve-bar-chart>
                 </div>
                </div>
  </div>
</div>
</template>

<script>
import Vue from 'vue';
import { VeBarChart } from 've-charts';


export default{
    name:'Comm',
    data(){
        return {
            comm:[],
            calls:[],
            usec:[],
            usecPerCall:[],
            chartData:{}
        };
    },
    methods:{
    },
    components:{
        VeBarChart
    },
    created(){
         var me = this; 
        this.chartSettings = {
            direction:'row'
        },
         this.textStyle = {
        color: 'white'
      }
        Vue.axios.get("/data").then((response)=>(
       response.data.result["comm"].forEach(item=>{
               me.comm.push(item['comm']);
               me.calls.push(item['calls']);
               me.usec.push(item['usec']);
               me.usecPerCall.push(item['usec_per_call']);
           }),
          me.chartData={
            dimensions:{
                name:'command',
                data:me.comm.reverse()
            },
            measures:[
                {
                    name:'usec_per_call',
                    data: me.usecPerCall.reverse()
                },
                {
                    name:'calls',
                    data: me.calls.reverse()
                },
                {
                    name:'usec',
                    data:me.usec.reverse()
                }
            ]
        } 
            ));

    }
}
</script>


<style scoped>
.container{
    margin-top:10px;
    width:100%;
    height:200px;
    display:inline-block;
}
.container h3 {
    display:block;
    height:96px;
    line-height:96px;
    margin:0;
    padding:0;
}

.box{
    display:flex;
}
.box-item{
    padding-top:10px;
    background-color:#262334;
    margin:2px;
    width:98%;
}

.box-item h4 {
    display:block;
    height:30px;
    line-height:30px;
    margin:0;
    padding:0;
}

.box-item p{
    margin-left:10px;
} 

</style>