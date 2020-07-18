import java.io.BufferedReader;
import java.io.InputStreamReader;
import java.io.IOException;

public class Main {
  public static void main(String[] args) throws IOException {
    BufferedReader input = new BufferedReader( new InputStreamReader( System.in ) );
    
    Double noteA = Double.parseDouble( input.readLine() );
    Double noteB = Double.parseDouble( input.readLine() );
    Double noteC = Double.parseDouble( input.readLine() );

    Double weightA = 2.0;
    Double weightB = 3.0;
    Double weightC = 5.0;
    Double sumOfWeights = weightA + weightB + weightC;

    Double ponderedNoteA = noteA * weightA;
    Double ponderedNoteB = noteB * weightB;
    Double ponderedNoteC = noteC * weightC;

    Double ponderedMean = (ponderedNoteA + ponderedNoteB + ponderedNoteC) / sumOfWeights;

    System.out.printf("MEDIA = %.1f\n", ponderedMean);

  }
}