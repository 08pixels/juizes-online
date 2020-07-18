import java.io.BufferedReader;
import java.io.InputStreamReader;
import java.io.IOException;

public class Main {
  public static void main(String[] args) throws IOException {
    BufferedReader stdin = new BufferedReader( new InputStreamReader( System.in ));

    String linePeace1 = stdin.readLine();
    String linePeace2 = stdin.readLine();

    String []splitedLine1 = linePeace1.trim().split("\\s+");
    String []splitedLine2 = linePeace2.trim().split("\\s+");

    int idPeaceA       = Integer.parseInt(splitedLine1[0]);
    int amountPeaceA   = Integer.parseInt(splitedLine1[1]);
    double pricePeaceA = Double.parseDouble(splitedLine1[2]);

    int idPeaceB       = Integer.parseInt(splitedLine2[0]);
    int amountPeaceB   = Integer.parseInt(splitedLine2[1]);
    double pricePeaceB = Double.parseDouble(splitedLine2[2]);

    double totalValueA = amountPeaceA * pricePeaceA;
    double totalValueB = amountPeaceB * pricePeaceB;
    
    System.out.printf("VALOR A PAGAR: R$ %.2f\n", totalValueA + totalValueB);
  }
}