import java.io.BufferedReader;
import java.io.InputStreamReader;
import java.io.IOException;

public class Main {
  public static void main(String[] args) throws IOException {
    BufferedReader input = new BufferedReader( new InputStreamReader( System.in ) );

    Double numberA = Double.parseDouble( input.readLine() );
    Double numberB = Double.parseDouble( input.readLine() );

    Double weightA = 3.5;
    Double weightB = 7.5;
    Double ponderedMean = (numberA*weightA + numberB*weightB) / (weightA+weightB);

    System.out.printf("MEDIA = %.5f\n", ponderedMean);
  }
}
