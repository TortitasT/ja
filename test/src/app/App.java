package app;

import java.sql.Connection;
import java.sql.DriverManager;
import java.util.Scanner;

import com.google.common.base.Joiner;

import app.other.Other;

class App {
  public static void main(String[] args) {
    System.out.println("It Works!");
    Other.run();

    Joiner joiner = Joiner.on("; ").skipNulls();
    var test = joiner.join("Harry", null, "Ron", "Hermione");

    System.out.println(test);

    try {
      Connection connection = DriverManager
          .getConnection("jdbc:mariadb://localhost:3306/dam2?user=root&password=password");

      System.out.println("Connected!");

      Scanner scanner = new Scanner(System.in);
      scanner.next();

      connection.close();
      scanner.close();
    } catch (Exception e) {
      System.out.println("Error: " + e.getMessage());
    }
  }
}
