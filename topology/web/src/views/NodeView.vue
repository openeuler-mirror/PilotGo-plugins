<template>
  <h1 class="h1">单机拓扑图演示页面</h1>
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
    const data = await topo.single_host_tree("d6900229-971b-482c-b1a5-cf108777c6db");
    console.log(data.data.tree);

    let root = data.data.tree
    root.img = server_logo;
    root.type = "image";
    root.size = 40

    initGraph(data.data.tree);
  } catch (error) {
    console.error(error)
  }
})

function initGraph(data: any) {
  let graph = new G6.TreeGraph({
    container: "topo-container",
    width: document.getElementById("topo-container")!.clientWidth,
    height: document.getElementById("topo-container")!.clientHeight,
    modes: {
      default: ['drag-canvas', 'zoom-canvas', "click-select", "drag-node",
        {
          type: 'collapse-expand',
          onChange: function onChange(item: any, collapsed) {
            const data = item.getModel();
            data.collapsed = collapsed;
            return true;
          },
        },
      ],
    },
    layout: {
      type: 'dendrogram',
      direction: 'LR',
      nodeSep: 30,
      rankSep: 100,
    },
  });
  graph.node(function (node: any) {
    // console.log(node);
    return {
      label: node.node.type + ":" + node.node.name,
      labelCfg: {
        position: node.children && node.children.length > 0 ? 'left' : 'right',
        offset: 5,
      },
    };
  });
  graph.on("nodeselectchange", (e) => {
    if (e.select) {
      let node = (e.target as any)._cfg!;
      let node_id = node.id
      console.log("click node:", node_id, e);

      title.value = "I am " + node_id;
      drawer.value = drawer.value ? false : true;
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
