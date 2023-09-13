<template>
  <h1 class="h1">集群拓扑图演示页面</h1>
  <div id="topo-container" class="container"></div>
  <el-drawer class="drawer" v-model="drawer" :title="title" :direction="direction" :before-close="handleClose">
    <span>Hi, there!</span>
  </el-drawer>
</template>

<script setup lang="ts">
import G6 from '@antv/g6';
import { ref, onMounted } from "vue";
import { topo } from '../request/api';
import server_logo from "@/assets/icon/server.png";

let drawer = ref(false)
const direction = 'rtl'
const title = ref("")

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
      let node_id = (e.target as any)._cfg!.id
      console.log("click node:", node_id);

      title.value = "I am " + node_id;
      drawer.value = drawer.value ? false : true;
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
