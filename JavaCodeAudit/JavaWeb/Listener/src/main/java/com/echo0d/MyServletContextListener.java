package com.echo0d;

import javax.servlet.ServletContextEvent;
import javax.servlet.ServletContextListener;
import javax.servlet.annotation.WebListener;

@WebListener
public class MyServletContextListener implements ServletContextListener {

    @Override
    public void contextInitialized(ServletContextEvent sce) {
        // 应用程序启动时执行的初始化操作
        System.out.println("应用程序已启动");
        // 加载配置文件、初始化数据库连接池等
        // ...
    }

    @Override
    public void contextDestroyed(ServletContextEvent sce) {
        // 应用程序关闭时执行的清理操作
        System.out.println("应用程序即将关闭");
        // 关闭数据库连接池、释放资源等
        // ...
    }
}