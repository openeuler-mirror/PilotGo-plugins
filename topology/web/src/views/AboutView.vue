<template>
  <h1 class="h1">拓扑图演示页面</h1>
  <div id="topo-container" class="container"></div>
</template>

<script setup lang="ts">
import G6 from '@antv/g6';
import { onMounted } from "vue";

onMounted(() => {
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

  let width = document.getElementById("topo-container")!.clientWidth;
  let height = document.getElementById("topo-container")!.clientHeight;
  console.log("canvas width: " + width + " height: " + height)

  let graph = new G6.Graph({
    container: "topo-container",
    width: document.getElementById("topo-container")!.clientWidth,
    height: document.getElementById("topo-container")!.clientHeight,
    modes: {
      default: ['drag-canvas', 'zoom-canvas', 'drag-node'],
    }
  });
  graph.data(data);
  graph.render();
  graph.fitCenter();

  window.onresize = () => {
    let width = document.getElementById("topo-container")!.clientWidth;
    let height = document.getElementById("topo-container")!.clientHeight;
    graph.changeSize(
      document.getElementById("topo-container")!.clientWidth,
      document.getElementById("topo-container")!.clientHeight)
    graph.fitCenter()
    console.log("resize: " + width + "," + height);
  }

})

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
</style>
