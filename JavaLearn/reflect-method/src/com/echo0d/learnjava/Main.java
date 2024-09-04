import java.lang.reflect.Method;

public class Main {
    public static void main(String[] args) throws Exception {
        // 通过Class实例获取所有Method信息。
        Class stdClass = Student.class;
        // 获取public方法getScore，参数为String:
        System.out.println(stdClass.getMethod("getScore", String.class));
        // 获取继承的public方法getName，无参数:
        System.out.println(stdClass.getMethod("getName"));
        // 获取private方法getGrade，参数为int:
        System.out.println(stdClass.getDeclaredMethod("getGrade", int.class));

        // 反射调用方法
        Student std = new Student();
        Method getGrade = std.getClass().getDeclaredMethod("getGrade", int.class);
        getGrade.setAccessible(true);
        Integer grade = (Integer) getGrade.invoke(std, 5);
        System.out.println(grade);

        // 直接调用方法private方法,会报错
        // System.out.println(std.getScore(2));

        // 多态
        // 获取Person的hello方法:
        Method h = Person.class.getMethod("hello");
        // 对Student实例调用hello方法:
        h.invoke(new Student());
        // 相当于
        Person p = new Student();
        p.hello();
    }
}
