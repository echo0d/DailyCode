/**
 * @author : echo0d
 * @date : 2024/8/15 20:34
 * @Description :
 */
import groovy.lang.GroovyShell;

public class groovyShellDemo {
    public static void main(String[] args) {
        GroovyShell shell = new GroovyShell();
        String cmd = "\"whoami\".execute().text";
        System.out.println(shell.evaluate(cmd));
    }
}
