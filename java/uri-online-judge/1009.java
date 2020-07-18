import java.io.IOException;
import java.io.BufferedReader;
import java.io.InputStreamReader;
import java.text.DecimalFormat;

public class Main {
  public static void main(String[] args) throws IOException {
    BufferedReader stdin = new BufferedReader( new InputStreamReader( System.in ) );
    DecimalFormat dc = new DecimalFormat("0.00");

    String name      = stdin.readLine();
    double salary    = Double.parseDouble( stdin.readLine() );
    double totalSold = Double.parseDouble( stdin.readLine() );

    double commissionRate = 0.15;
    double commission = totalSold * commissionRate;

    double totalSalary = salary + commission;

    System.out.println("TOTAL = R$ " + dc.format(totalSalary));
  }
} 