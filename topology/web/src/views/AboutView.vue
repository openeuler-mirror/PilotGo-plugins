<template>
  <h1 class="h1">拓扑图演示页面</h1>
  <div id="topo-container" class="container"></div>

  <el-drawer class="drawer" v-model="drawer" :title="title" :direction="direction" :before-close="handleClose">
    <span>Hi, there!</span>
  </el-drawer>
</template>

<script setup lang="ts">
import G6 from '@antv/g6';
import { ref, onMounted } from "vue";


let drawer = ref(false)
const direction = 'rtl'
const title = ref("")

function handleClose() {
 drawer.value = false
}

onMounted(() => {
  initGraph()

})

function initGraph() {
  const data = {
    canvasWidth: 0,
    canvasHeight: 0,

    // 节点
    nodes: [
      {
        id: 'node1',
        x: 100,
        y: 200,
      },
      {
        id: 'node2',
        x: 300,
        y: 200,
      },
    ],
    // 边集
    edges: [
      // 表示一条从 node1 节点连接到 node2 节点的边
      {
        source: 'node1',
        target: 'node2',
      },
    ],
  };

  let graph = new G6.Graph({
    container: "topo-container",
    width: document.getElementById("topo-container")!.clientWidth,
    height: document.getElementById("topo-container")!.clientHeight,
    modes: {
      default: ['drag-canvas', 'zoom-canvas', 'drag-node', 'click-select'],
    }
  });
  graph.on("nodeselectchange", (e) => {
    if (e.select) {
      let node_id = e.target._cfg!.id
      console.log("click node:", node_id);

      title.value = "I am " + node_id;
      drawer.value = drawer.value?false:true;
    } else {
      console.log("node unselected")
    }
    return false
  });
  graph.data(data);
  graph.render();
  graph.fitCenter();

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
