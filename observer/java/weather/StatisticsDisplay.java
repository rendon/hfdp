public class StatisticsDisplay implements Observer, DisplayElement {
    private float maxTemp = 0.0f;
    private float minTemp = 200;
    private float tempSum = 0.0f;
    private int numReadings;
    private Subject weatherData;

    public StatisticsDisplay(Subject weatherData) {
        this.weatherData = weatherData;
        weatherData.registerObserver(this);
    }

    public void update(float t, float h, float p) {
        tempSum += t;
        if (t > maxTemp) {
            maxTemp = t;
        }

        if (t < minTemp) {
            minTemp = t;
        }
        display();
    }

    public void display() {
        System.out.print("Avg/Max/Min temperature = ");
        System.out.print(tempSum / numReadings);
        System.out.println(" / " + maxTemp + "/" + minTemp);
    }
}
