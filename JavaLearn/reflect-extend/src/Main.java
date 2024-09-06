// reflection
public class Main {
    @SuppressWarnings("rawtypes")
    public static void main(String[] args) throws Exception {
        // 获取父类的Class
        System.out.println("获取父类的Class");
        Class i = Integer.class;
        Class n = i.getSuperclass();
        System.out.println(n);
        Class o = n.getSuperclass();
        System.out.println(o);
        System.out.println(o.getSuperclass());

        // 获取Interface
        System.out.println("获取Interface");
        Class[] is = i.getInterfaces();
        for (Class c : is) {
            System.out.println(c);
        }
        // 获取父类的Interface
        System.out.println("获取父类的Interface");
        Class[] ns = n.getInterfaces();
        for (Class c : ns) {
            System.out.println(c);
        }
        // java.io.FilterInputStream，因为DataInputStream继承自FilterInputStream
        System.out.println(java.io.DataInputStream.class.getSuperclass()); 

        // null，对接口调用getSuperclass()总是返回null，获取接口的父接口要用getInterfaces()
        System.out.println(java.io.Closeable.class.getSuperclass()); 

        CharSequence cs = new StringBuilder();
    }
}