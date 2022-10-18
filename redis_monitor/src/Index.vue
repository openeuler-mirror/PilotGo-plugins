<template>
<div>
<div  class="first">
服务器版本:{{server['os']}}
</div>

<div class="container">
       <!--     <h3 title="flex-direction:row为横向,row-reverse为横向倒序" align="center">Flex Item Order</h3> -->
            <div class="box" style="flex-direction: row;">
                <div class="box-item" style="order: 1;flex-grow: 1;height: 188px;margin-left:10px;">
                <h4 align="center">系统</h4>
                <div>
                <p>主机操作系统:{{server['os']}}</p>
                 <p>主机架构：{{server['arch_bits']}}</p>
                 <p>监听端口：{{server['tcp_port']}}</p>
                 </div>
                </div>


                <div class="box-item" style="order: 2;flex-grow: 1;height: 188px;">
                <h4 align="center">内存</h4>
                <div>
                <p>主机操作系统:{{server['os']}}</p>
                 <p>主机架构：{{server['arch_bits']}}</p>
                 <p>监听端口：{{server['tcp_port']}}</p>
                 </div>
                </div>


                <div class="box-item" style="order: 3;flex-grow: 1;height: 188px;margin-right:10px;">
                                <h4 align="center">CPU</h4>
                <div>
                <p>消耗的系统CPU:{{cpu['used_cpu_sys']}}</p>
                 <p>消耗的用户CPU:{{cpu['used_cpu_user']}}</p>
                 <p>后台进程占用的系统CPU:{{cpu['used_cpu_sys_children']}}</p>
                 <p>后台进程占用的用户CPU:{{cpu['used_cpu_user_children']}}</p>
                 </div>
                </div>
            </div>
        </div>

<Comm></Comm>


<div class="container"  style="margin-top:20px;height:500px;background-color:#262334;">

    <h4 align="center" style="    display:block;
    height:30px;
    line-height:30px;color:#a4a4a4">统计信息</h4>
   <div class="box" style="flex-direction: row;">
                <div class="box-item" style="order: 3;flex-grow: 1;height: 188px;">
               <ve-polar-chart :data="chartData1" :settings="settings1" />
                </div>

                <div class="box-item" style="order: 3;flex-grow: 1;height: 188px;">
               <ve-polar-chart :data="chartData2" :settings="settings2" />
                </div>


                                <div class="box-item" style="order: 3;flex-grow: 1;height: 188px;">
               <ve-polar-chart :data="chartData3" :settings="settings3" />
                </div>

   </div>

</div>

</div>
</template>
<script>
import Vue from 'vue';
import Comm from './components/Comm';
import { VePolarChart } from 've-charts';

export default{
    name:'Index',
    components:{
        Comm,
        VePolarChart
    },
    data(){
        return {
            server:[],
            cpu:[],
            charsData1:[],
            charsData2:[],
            charsData3:[],
            chartData1:{},
            chartData2:{},
            chartData3:{}
    }
    },
    created(){}
}
</script>

<style>
.first{
    text-align:center;
    padding-top:50px;
    background:#14121f;
    color:azure;
    min-height:100px;
    position:relative;
    color:#a4a4a4;
}
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
    width:95%;
    border-radius: 8px;
}
.box-item p {
    color:#a4a4a4;
}
.box-item h4 {
    display:block;
    height:30px;
    line-height:30px;
    margin:0;
    padding:0;
        color:#a4a4a4;
}

.box-item p{
    margin-left:10px;
} 


</style>