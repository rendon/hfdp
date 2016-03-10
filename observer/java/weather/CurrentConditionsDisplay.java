public class CurrentConditionsDisplay implements Observer, DisplayElement {
    private float temperature;
    private float humidity;
    private Subject weatherData;

    public CurrentConditionsDisplay(Subject weatherData) {
        this.weatherData = weatherData;
        this.weatherData.registerObserver(this);
    }

    public void update(float t, float h, float p) {
        temperature = t;
        humidity = h;
        display();
    }

    public void display() {
        System.out.print("Current conditions: " + temperature + "F degrees");
        System.out.println(" and " + humidity + "% humidity");
    }
}
