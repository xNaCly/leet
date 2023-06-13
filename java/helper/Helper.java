package helper;

public class Helper {
    public static void Assert(Object a, Object b) {
        if (a != b)
            throw new AssertionError(String.format("%s not equal to %s", a, b));
    }
}
