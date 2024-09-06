
public class Main {
    public static void main(String[] args) throws ReflectiveOperationException {
        Person person = new Person("Xiao Ming", "北京");
        person.check(person); // 调用check方法，检查person对象是否合法
        System.out.println(person.city); // 输出person对象的字符串表示
    }
   
}
