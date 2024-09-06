package com.echo0d.proxy;
import java.lang.reflect.InvocationHandler;

public class HelloDynamicProxy implements Hello {
    InvocationHandler handler;
    public HelloDynamicProxy(InvocationHandler handler) {
        this.handler = handler;
    }
    public void morning(String name){
        try {
            handler.invoke(
               this,
               Hello.class.getMethod("morning", String.class),
               new Object[] { name }
            );
        } catch (Throwable e) {
            e.printStackTrace();
        }
    }
}