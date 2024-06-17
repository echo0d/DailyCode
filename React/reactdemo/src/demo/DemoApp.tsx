import React from 'react';
import { Input, Button, Row, Col } from 'antd';

//引入刚刚创建的子组件列表，以便在后面使用<TaskList></TaskList>组件
import TaskList from './TaskList';
//引入DemoFuncs.ts文件
import { DataInfoType, taskNameChange, addNewTaskToList, checkedOrNotchecked, deleteCheckedTaskName} from './DemoFuncs';

function DemoApp() {

  //以泛型方式定义state的数据，名初始化state中的业务数据对象
  const [dataInfo, setDataInfo] = React.useState<DataInfoType>({
    taskName: '',
    taskList: []
});

  return (
    <React.Fragment>
      <h1 style={{ textAlign: "center" }}>待办列表</h1>
      <hr></hr>
      <Row align="middle" gutter={5}>
        <Col span={1}>任务标题</Col>
        <Col span={11}>
          {/** value属性与state对象中的taskName绑定 */}
          {/** 增加onChange事件处理 */}
          <Input
            value={dataInfo.taskName}
            onChange={(e) => taskNameChange(e.target.value, dataInfo, setDataInfo)}
          ></Input>
        </Col>
        <Col span={12}>
          {/** 增加onClick事件处理 */}
          <Button type="primary" onClick={() => addNewTaskToList(dataInfo, setDataInfo)}>
            添加
          </Button>
        </Col>
      </Row>
      <Row>
        <Col span={24}>
          <ul>
            {/** 使用子组件显示待办任务列表 */}
            <TaskList taskDataInfo={dataInfo} checkboxChange={(e) => checkedOrNotchecked(e, dataInfo, setDataInfo)}></TaskList>
          </ul>
        </Col>
      </Row>
      <Row gutter={5}>
        <Col span={12}></Col>
        <Col span={12}>
          <Button type="primary" onClick={() => deleteCheckedTaskName(dataInfo,setDataInfo)}>删除</Button>
        </Col>
      </Row>
    </React.Fragment>
  );
}

export default DemoApp;
