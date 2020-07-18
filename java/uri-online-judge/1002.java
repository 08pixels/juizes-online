import java.io.BufferedReader;
import java.io.InputStreamReader;
import java.io.IOException;

public class Main {
  public static void main( String[] args ) throws IOException {
    BufferedReader br = new BufferedReader( new InputStreamReader( System.in ) );

    double radius = Double.parseDouble( br.readLine() );

    System.out.printf( "A=%.4f\n", 3.14159 * Math.pow( radius, 2 ) );
  }
}