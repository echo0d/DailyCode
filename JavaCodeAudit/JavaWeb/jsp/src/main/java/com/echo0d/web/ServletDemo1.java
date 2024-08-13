package com.echo0d.web;


import com.echo0d.pojo.User;

import javax.servlet.*;
import javax.servlet.http.*;
import javax.servlet.annotation.*;
import java.io.IOException;
import java.util.ArrayList;
import java.util.List;

@WebServlet("/demo1")
public class ServletDemo1 extends HttpServlet {
    @Override
    protected void doGet(HttpServletRequest request, HttpServletResponse response) throws ServletException, IOException {
        //1. 准备数据
        List<User> Users = new ArrayList<User>();
        Users.add(new User("uName01","man","123456"));
        Users.add(new User("uName02","woman","123456789"));
        Users.add(new User("uName03","man","123456"));
        //2. 存储到request域中
        request.setAttribute("Users",Users);
        request.setAttribute("status",1);
        //3. 转发
        request.getRequestDispatcher("/el-demo.jsp").forward(request,response);
    }

    @Override
    protected void doPost(HttpServletRequest request, HttpServletResponse response) throws ServletException, IOException {
        this.doGet(request, response);
    }
}
