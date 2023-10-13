package app;

import java.sql.Connection;
import java.sql.DriverManager;
import java.util.Scanner;

import app.other.Other;

class App {
  public static void main(String[] args) {
    System.out.println("la ignorancia hace la felicidad");
    Other.mojon();

    try {
      Connection connection = DriverManager
          .getConnection("jdbc:mariadb://localhost:3306/dam2?user=root&password=password");

      System.out.println("Conexión realizada con éxito");

      Scanner scanner = new Scanner(System.in);
      scanner.next();
    } catch (Exception e) {
      System.out.println("Error: " + e.getMessage());
    }
  }
}
