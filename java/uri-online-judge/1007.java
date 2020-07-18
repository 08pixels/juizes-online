import java.io.IOException;
import java.io.BufferedReader;
import java.io.InputStreamReader;

public class Main {
  public static void main(String []args) throws IOException {
    BufferedReader stdin = new BufferedReader( new InputStreamReader( System.in ));
    
    int numberA = Integer.parseInt( stdin.readLine() );
    int numberB = Integer.parseInt( stdin.readLine() );
    int numberC = Integer.parseInt( stdin.readLine() );
    int numberD = Integer.parseInt( stdin.readLine() );

    int productAB = numberA * numberB;
    int productCD = numberC * numberD;

    int difference = productAB - productCD;

    System.out.printf("DIFERENCA = %d\n", difference);

  }
}