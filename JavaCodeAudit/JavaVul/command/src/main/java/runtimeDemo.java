
import java.io.BufferedReader;
import java.io.InputStreamReader;

/**
 * @author : echo0d
 * @date : 2024/8/15 20:24
 * @Description :
 */
public class runtimeDemo {
    public static void main(String[] args) {
        try {
            Process process = Runtime.getRuntime().exec("calc");
            BufferedReader reader = new BufferedReader(new InputStreamReader(process.getInputStream()));
            String line;
            while ((line = reader.readLine()) != null) {
                System.out.println(line);
            }
        } catch (Exception e) {
            e.printStackTrace();
        }

    }
}
