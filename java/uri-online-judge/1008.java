import java.io.BufferedReader;
import java.io.InputStreamReader;
import java.io.IOException;

public class Main {
  public static void main(String []args) throws IOException {
    BufferedReader stdin = new BufferedReader( new InputStreamReader( System.in ));

    int id = Integer.parseInt( stdin.readLine() );
    int workedHours = Integer.parseInt( stdin.readLine() );
    double hourlyPrice = Double.parseDouble( stdin.readLine() );

    double salary = workedHours * hourlyPrice;

    System.out.printf("NUMBER = %d\n", id);
    System.out.printf("SALARY = U$ %.2f\n", salary);
  }
}
