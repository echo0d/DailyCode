public class Person {
    String name;

    String getName() {
        return this.name;
    }
    void setName(String name) {
        this.name = name;
    }
    private Person(String name) {
        this.name = name;
    }
}
