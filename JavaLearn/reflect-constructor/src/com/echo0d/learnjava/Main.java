import java.lang.reflect.Constructor;

public class Main {
    @SuppressWarnings("rawtypes")
    public static void main(String[] args) throws Exception {
        // 获取构造方法Integer(int):
        Constructor cons1 = Integer.class.getConstructor(int.class);
        // 调用构造方法:
        Integer n1 = (Integer) cons1.newInstance(123);
        System.out.println(n1);

        // 获取构造方法Integer(String)
        Constructor cons2 = Integer.class.getConstructor(String.class);
        Integer n2 = (Integer) cons2.newInstance("456");
        System.out.println(n2);

        // 获取Private的构造方法
        Constructor cons3 = Person.class.getDeclaredConstructor(String.class);
        cons3.setAccessible(true);
        Person p = (Person) cons3.newInstance("Bob");
        System.out.println(p.name);
    }

}
