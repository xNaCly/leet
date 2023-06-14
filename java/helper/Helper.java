package helper;

import java.util.Arrays;
import java.util.stream.Collectors;

public class Helper {
    public static void Assert(Object a, Object b) {
        if (a != b)
            throw new AssertionError(String.format("%s not equal to %s", a, b));
    }

    public static void AssertArrayEquals(Object[] a, Object[] b) {
        boolean r = false;
        if (a.length != b.length) {
            r = true;
        }

        for (int i = 0; i < a.length; i++) {
            if (a[i] != b[i]) {
                r = true;
                break;
            }
        }

        if (r) {
            throw new AssertionError(
                    String.format("{%s} not equal to {%s}",
                            Arrays.stream(a).map(i -> i.toString()).collect(Collectors.joining(", ")),
                            Arrays.stream(b).map(i -> i.toString()).collect(Collectors.joining(", "))));
        }
    }
}
