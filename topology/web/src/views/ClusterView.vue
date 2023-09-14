<template>
  <h1 class="h1">集群拓扑图演示页面</h1>
  <div id="topo-container" class="container"></div>
  <el-drawer class="drawer" v-model="drawer" :title="title" direction="rtl" :before-close="handleClose">
    <el-table :data="table_data" stripe style="width: 100%">
      <el-table-column prop="name" label="属性" width="180" />
      <el-table-column prop="value" label="值"/>
    </el-table>
  </el-drawer>
</template>

<script setup lang="ts">
import G6 from '@antv/g6';
import { ref, reactive, onMounted } from "vue";
import { topo } from '../request/api';
import server_logo from "@/assets/icon/server.png";

let drawer = ref(false)
let title = ref("")
let table_data = reactive<any>([])

function handleClose() {
  drawer.value = false
}

onMounted(async () => {
  try {
    const data = await topo.multi_host_topo();
    // console.log(data.data);

    for (let i = 0; i < data.data.nodes.length; i++) {
      let node = data.data.nodes[i];
      if (node.type === "host") {
        node.img = server_logo;
        node.type = "image";
        node.size = 40;
        let ip = node.id.split("_").pop()
        node.label = ip;
      } else if (node.type === "process") {
        node.label = node.name + ":" + node.metrics.Pid;
      } else if (node.type === "net") {
        node.label = node.name;
      }
    };

    initGraph(data.data);
  } catch (error) {
    console.error(error)
  }
})

function initGraph(data: any) {
  let graph = new G6.Graph({
    container: "topo-container",
    width: document.getElementById("topo-container")!.clientWidth,
    height: document.getElementById("topo-container")!.clientHeight,
    layout: {
      type: 'force',
      preventOverlap: true,
      linkDistance: 100,
    },
    modes: {
      default: ['drag-canvas', 'zoom-canvas', "click-select", "drag-node"],
    },
  });
  graph.node(function (node) {
    return {
      labelCfg: {
        position: "bottom",
        offset: 5,
      },
    };
  });
  graph.on("nodeselectchange", (e) => {
    if (e.select) {
      let node = (e.target as any)._cfg
      console.log("click node:", node.id);

      updateDrawer(node)
    } else {
      console.log("node unselected")
    }
    return false
  });
  graph.data(data);
  graph.render();

  window.onresize = () => {
    graph.changeSize(
      document.getElementById("topo-container")!.clientWidth,
      document.getElementById("topo-container")!.clientHeight)
    graph.fitCenter()
  }
}

function updateDrawer(node: any) {
  title.value = node.id + "节点属性";
  drawer.value = drawer.value ? false : true;

  // console.log(node)
  let metrics = node.model.metrics;
  for (let key in metrics) {
    table_data.push({
      name: key,
      value: metrics[key],
    })
  };
}

</script>

<style scoped>
.h1 {
  width: 100%;
  text-align: center;
}

.container {
  width: 100%;
  height: 100%;
  background-color: white;
}

.drawer {
  height: 100%;
}
</style>
