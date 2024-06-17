import React from 'react';
import ReactDOM from 'react-dom/client';
import { RouterProvider } from 'react-router-dom';
import {routes} from './routers/RoutesDef';
<<<<<<< HEAD
=======
import 'antd/dist/reset.css';
>>>>>>> 88f3721c8f912a5adef2504963fecd862afd0000

const root = ReactDOM.createRoot(
  document.getElementById('root') as HTMLElement
);
root.render(
  <React.StrictMode>
    <RouterProvider router={routes}></RouterProvider>{/** 使用路由定义数据 */}
  </React.StrictMode>
);
