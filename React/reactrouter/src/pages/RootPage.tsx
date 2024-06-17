import React from "react";
import { Link } from "react-router-dom";
import { Outlet } from "react-router-dom";

function RootPage() {
  return (
    <React.Fragment>
      <div style={{ float: "left", width: "200px", height: "600px" }}>
        左侧功能区
        <ul>
          <li>
            <Link to={"/one"}>功能一</Link>
          </li>
          {/** 修改to属性的值，123就是点击链接是传递的业务数据的值 */}
          <li>
            <Link to={"/two/123"}>功能二</Link>
          </li>
        </ul>
      </div>
      <div style={{ height: "600px" }}>
        右侧主操作区
        <Outlet />
      </div>
    </React.Fragment>
  );
}

export default RootPage;
