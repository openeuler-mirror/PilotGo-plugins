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

let drawer = ref(false)
const direction = 'rtl'
const title = ref("")

function handleClose() {
  drawer.value = false
}

onMounted(async () => {
  try {
    const data = await topo.single_host_topo("8140418e-dd3e-44f6-b769-1b2b9e562a3e");
    console.log(data.data);
    let d = data.data

    initGraph(parse(d));
  } catch (error) {
    console.error(error)
  }
})

function parse(data: any) {
  let root = {
    id: "8140418e-dd3e-44f6-b769-1b2b9e562a3e",
    children: [
      {
        id: "process",
        children: [] as any[],
      },
      {
        id: "resource",
        children: [] as any[],
      },
    ],
  }

  for (const node of data.nodes) {
    const words: string[] = node.id.split("_");
    // let id = words[0];
    let resource_type = words[1];
    let resource = words[2];

    if (resource_type == "process") {
      root.children[0].children.push(node)
    } else if (resource_type == "resource") {
      root.children[1].children.push(node)
    }
  }

  // for (const node of data.nodes) {
  //   const words:string[] = node.id.split("_");
  //   let id = words[0];
  //   let resource_type = words[1];
  //   let resource = words[2];
  //   if (resource_type == "net") {

  //   } else if (resource_type == "thread") {

  //   }
  // }

  return root
}

function initGraph(data: any) {
  // const data = {
  //   "id": "Modeling Methods",
  //   "children": [
  //     {
  //       "id": "Classification",
  //       "children": [
  //         {
  //           "id": "Logistic regression"
  //         },
  //         {
  //           "id": "Linear discriminant analysis"
  //         },
  //         {
  //           "id": "Rules"
  //         },
  //         {
  //           "id": "Decision trees"
  //         },
  //         {
  //           "id": "Naive Bayes"
  //         },
  //         {
  //           "id": "K nearest neighbor"
  //         },
  //         {
  //           "id": "Probabilistic neural network"
  //         },
  //         {
  //           "id": "Support vector machine"
  //         }
  //       ]
  //     },
  //     {
  //       "id": "Consensus",
  //       "children": [
  //         {
  //           "id": "Models diversity",
  //           "children": [
  //             {
  //               "id": "Different initializations"
  //             },
  //             {
  //               "id": "Different parameter choices"
  //             },
  //             {
  //               "id": "Different architectures"
  //             },
  //             {
  //               "id": "Different modeling methods"
  //             },
  //             {
  //               "id": "Different training sets"
  //             },
  //             {
  //               "id": "Different feature sets"
  //             }
  //           ]
  //         },
  //         {
  //           "id": "Methods",
  //           "children": [
  //             {
  //               "id": "Classifier selection"
  //             },
  //             {
  //               "id": "Classifier fusion"
  //             }
  //           ]
  //         },
  //         {
  //           "id": "Common",
  //           "children": [
  //             {
  //               "id": "Bagging"
  //             },
  //             {
  //               "id": "Boosting"
  //             },
  //             {
  //               "id": "AdaBoost"
  //             }
  //           ]
  //         }
  //       ]
  //     },
  //     {
  //       "id": "Regression",
  //       "children": [
  //         {
  //           "id": "Multiple linear regression"
  //         },
  //         {
  //           "id": "Partial least squares"
  //         },
  //         {
  //           "id": "Multi-layer feedforward neural network"
  //         },
  //         {
  //           "id": "General regression neural network"
  //         },
  //         {
  //           "id": "Support vector regression"
  //         }
  //       ]
  //     }
  //   ]
  // }

  // let graph = new G6.Graph({
  let graph = new G6.TreeGraph({
    container: "topo-container",
    width: document.getElementById("topo-container")!.clientWidth,
    height: document.getElementById("topo-container")!.clientHeight,
    modes: {
      default: ['drag-canvas', 'zoom-canvas', "click-select", "drag-node",
      // default: ['drag-canvas', 'zoom-canvas',
        {
          type: 'collapse-expand',
          onChange: function onChange(item, collapsed) {
            const data = item.getModel();
            data.collapsed = collapsed;
            return true;
          },
        },
      ],
    },
    layout: {
      type: 'dendrogram',
      direction: 'LR', // H / V / LR / RL / TB / BT
      nodeSep: 30,
      rankSep: 100,
    },
  });
  graph.node(function (node) {
    return {
      label: node.id,
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
