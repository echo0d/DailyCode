

public class NativeExample {
    static {
        System.loadLibrary("native_lib");
    }

    public native void sayHello();

    public static void main(String[] args) {
        new NativeExample().sayHello();
        System.out.println("Hello from Java!");
    }
}
