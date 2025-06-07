import { createRouter, createWebHistory } from "vue-router";
import SubjectsView from "../views/SubjectsView.vue";
import ClassesView from "../views/ClassesView.vue";
import ClassroomView from "../views/ClassroomView.vue";
import ClassCover from "../views/Class/ClassCover.vue";

const routes = [
  {
    path: "/",
    name: "Subjects",
    component: SubjectsView,
  },
  {
    path: "/class",
    name: "Classes",
    component: ClassesView,
  },
  {
    path: "/classroom",
    name: "Classroom",
    component: ClassroomView,
    children: [
      {
        path: "",
        name: "ClassCover",
        component: ClassCover,
      },
    ],
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
