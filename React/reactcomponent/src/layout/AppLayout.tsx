import { Layout, ConfigProvider, theme } from "antd";
import './AppLayout.css';//添加import css代码
import AppMenu from "./AppMenu";//引入菜单组件
import AppTabs from "./AppTabs";

const { Header, Sider, Content } = Layout;


function AppLayout() {
    return (
        <ConfigProvider theme={{ algorithm: theme.darkAlgorithm }}>{/** 使用antd的暗黑模式 */}
            <Layout>
                <Header>Header部分</Header>
                <Layout>
                    <Sider width={190} style={{overflow: "auto"}}>
                        <span style={{color: "white"}}>菜单区</span>
                        <AppMenu />{/** 在左侧边栏应用菜单组件 */}
                    </Sider>
                    <Content>
                        <span style={{color: "white"}}>页签区</span>
                        <AppTabs />{/** 在右侧内容区应用页签组件 */}
                    </Content>
                </Layout>
            </Layout>
        </ConfigProvider>
    );
};

export default AppLayout;