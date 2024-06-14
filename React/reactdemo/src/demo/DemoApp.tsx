import React from "react";
import { Input, Button, Row, Col, Checkbox } from "antd";

function DemoApp() {
  const [dataInfo, setDataInfo] = React.useState({
    taskName: "新任务名称",
    taskList: ["编写React教程", "吃饭", "睡觉"],
  });
  /** 定义函数onChange处理函数，调用setDataInfo函数修改state的值 */
  function taskNameChanged(newTaskname: string) {
    let newStateInfo = { ...dataInfo };
    newStateInfo.taskName = newTaskname;
    setDataInfo(newStateInfo);
  }

  /** 添加按钮处理函数 */
  function addNewTaskToList() {
    let newStateInfo = {...dataInfo};
    newStateInfo.taskList = [...newStateInfo.taskList, newStateInfo.taskName];
    setDataInfo(newStateInfo);
}
  
  return (
    <React.Fragment>
      <h1 style={{ textAlign: "center" }}>待办列表</h1>
      <hr></hr>
      <Row align="middle" gutter={5}>
        <Col span={1}>任务标题</Col>
        {/** value属性与state对象中的taskName绑定 */}
        {/* <Col span={11}><Input value={dataInfo.taskName}></Input></Col> */}
        {/** 增加onChange事件处理 */}
        <Col span={11}>
          <Input
            value={dataInfo.taskName}
            onChange={(e) => taskNameChanged(e.target.value)}
          ></Input>
        </Col>
        {/* <Col span={12}>
          <Button type="primary">添加</Button>
        </Col> */}
        {/** 增加onClick事件处理 */}
        <Col span={12}><Button type='primary' onClick={addNewTaskToList}>添加</Button></Col>
      </Row>
      <Row>
        <Col span={24}>
          <ul>
            {/** 将原来的静态内容修改为根据taskList内容输出列表 */}
            {dataInfo.taskList.map((oneItem) => (
              <li key={oneItem}>
                <Checkbox></Checkbox>
                {oneItem}
              </li>
            ))}
          </ul>
        </Col>
      </Row>
      <Row gutter={5}>
        <Col span={12}></Col>
        <Col span={12}>
          <Button type="primary">删除</Button>
        </Col>
      </Row>
    </React.Fragment>
  );
}

export default DemoApp;
