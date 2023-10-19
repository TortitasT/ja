package app;

import java.sql.Connection;
import java.sql.DriverManager;
import java.util.Scanner;

import org.apache.hc.client5.http.classic.methods.HttpGet;
import org.apache.hc.client5.http.impl.classic.CloseableHttpClient;
import org.apache.hc.client5.http.impl.classic.CloseableHttpResponse;
import org.apache.hc.client5.http.impl.classic.HttpClients;
import org.apache.hc.core5.http.HttpEntity;
import org.apache.hc.core5.http.ParseException;
import org.apache.hc.core5.http.io.entity.EntityUtils;

import app.other.Other;

class App {
  public static void main(String[] args) {
    System.out.println("It Works!");
    Other.mojon();

    String resultContent = null;
    HttpGet httpGet = new HttpGet("https://www.google.com");
    try (CloseableHttpClient httpclient = HttpClients.createDefault()) {
      try (CloseableHttpResponse response = httpclient.execute(httpGet)) {
        // Get status code
        System.out.println(response.getVersion()); // HTTP/1.1
        System.out.println(response.getCode()); // 200
        System.out.println(response.getReasonPhrase()); // OK
        HttpEntity entity = response.getEntity();
        // Get response information
        resultContent = EntityUtils.toString(entity);
      }
    } catch (IOException | ParseException e) {
      e.printStackTrace();
    }

    try {
      Connection connection = DriverManager
          .getConnection("jdbc:mariadb://localhost:3306/dam2?user=root&password=password");

      System.out.println("Conexión realizada con éxito");

      Scanner scanner = new Scanner(System.in);
      scanner.next();

      connection.close();
      scanner.close();
    } catch (Exception e) {
      System.out.println("Error: " + e.getMessage());
    }
  }
}
