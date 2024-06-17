import { createBrowserRouter } from "react-router-dom";
import RootPage from "../pages/RootPage"; //导入刚刚创建的组件
import PageOne from "../pages/PageOne";
import PageTwo, { loader as pageLoader } from "../pages/PageTwo"; //导入数据加载器

export const routes = createBrowserRouter([
  {
    path: "/", //路由path
    element: <RootPage />, //path根对应的组件，又import语句生成
    children: [
      {
        path: "/one",
        element: <PageOne />,
      },
      {
        path: "/two/:bizDataName", //冒号 ( : ) 具有特殊含义，将其转换为“动态段”，由<Link>组件定义变量bizDataName的值,
        element: <PageTwo />,
        loader: pageLoader, //定义数据加载器
      },
    ],
  },
]);

export default routes;
