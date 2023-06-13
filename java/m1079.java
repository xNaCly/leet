import java.util.HashMap;

import helper.Helper;

public class m1079 {
    public static void main(String[] args) {
        HashMap<String, Integer> test = new HashMap<>() {
            {
                put("AAB", 8);
                put("AAABBC", 188);
                put("V", 1);
            }
        };

        for (String key : test.keySet()) {
            Helper.Assert(numTilePossibilities(key), test.get(key));
        }
    }

    public static int numTilePossibilities(String tiles) {
        return 0;
    }
}
