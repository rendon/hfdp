import java.io.*;

public class InputTest {
    public static void main(String[] args) {
        try {
            InputStream in = new FileInputStream("test.txt");
            in = new BufferedInputStream(in);
            in = new LowerCaseInputStream(in);
            int c;
            while ((c = in.read()) >= 0) {
                System.out.print((char)c);
            }
        } catch (IOException e) {
            e.printStackTrace();
        }
    } 
}
