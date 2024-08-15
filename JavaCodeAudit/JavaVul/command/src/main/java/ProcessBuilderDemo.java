import java.io.BufferedReader;
import java.io.InputStreamReader;

/**
 * @author : echo0d
 * @date : 2024/8/15 20:26
 * @Description :
 */
public class ProcessBuilderDemo {
    public static void main(String[] args) {
        try {
            ProcessBuilder builder = new ProcessBuilder("ipconfig");
            Process process = builder.start();
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
