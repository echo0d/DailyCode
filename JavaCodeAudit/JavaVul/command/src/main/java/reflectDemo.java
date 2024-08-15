/**
 * @author : echo0d
 * @date : 2024/8/15 21:05
 * @Description :
 */
import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.lang.Class;
import java.lang.reflect.*;
public class reflectDemo {
    public static void main(String[] args) throws ClassNotFoundException, NoSuchMethodException, InvocationTargetException, IllegalAccessException {
        Class<?> c = Class.forName("java.lang.Runtime");//获取类
        Method m1 =  c.getMethod("getRuntime");//获取getRuntime方法，用于创建对象
        Method m2 = c.getMethod("exec",String.class);//获取exec方法，用于执行命令
        Object obj =  m1.invoke(null,null);//创建对象
        Process process = (Process) m2.invoke(obj,"whoami");//反射调用
        // 下面可以不要，直接m2.invoke(obj,"whoami"); 只是没回显
        try (BufferedReader reader = new BufferedReader(new InputStreamReader(process.getInputStream()))) {
            String line;
            while ((line = reader.readLine()) != null) {
                System.out.println(line);
            }
        } catch (IOException e) {
            throw new RuntimeException(e);
        }
    }
}
