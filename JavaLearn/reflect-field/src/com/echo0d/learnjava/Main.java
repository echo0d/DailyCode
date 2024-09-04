package com.echo0d.learnjava;

import java.lang.reflect.Field;
public class Main {

	public static void main(String[] args) throws Exception {
		String name = "Xiao Ming";
		int age = 20;
		Person p = new Person();
		Class c = p.getClass();
		// name
		Field fieldName = c.getDeclaredField("name");
		fieldName.setAccessible(true);
		fieldName.set(p, name);
		// age
		Field fieldAge = c.getDeclaredField("age");
		fieldAge.setAccessible(true);
		fieldAge.set(p, age);
		// 输出结果
		System.out.println(p.getName()); // "Xiao Hong"
		System.out.println(p.getAge()); // 15
	}
}
