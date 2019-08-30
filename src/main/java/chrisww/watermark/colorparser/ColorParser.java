package chrisww.watermark.colorparser;

import java.awt.*;
import java.util.HashMap;
import java.util.Map;

public class ColorParser {
    public static Color parse(String str) throws ParseError {
        str = str.toLowerCase();

        if (str.length() > 0 && str.charAt(0) == '#') {
            return parseHex(str);
        } else {
            return parseName(str);
        }
    }

    private static Color parseHex(String str) throws ParseError {
        if (str.length() != 7) {
            throw new ParseError("invalid color str");
        }

        int r = hexStrToInt(str.substring(1, 3));
        int g = hexStrToInt(str.substring(3, 5));
        int b = hexStrToInt(str.substring(5, 7));

        return new Color(r, g, b);
    }

    private static int hexStrToInt(String str) throws ParseError {
        int value = 0;
        for (int i = str.length() - 1; i >= 0; --i) {
            char c = str.charAt(i);
            int v = 0;
            if ('0' <= c && c <= '9') {
                v = c - '0';
            } else if ('a' <= c && c <= 'f') {
                v = c - 'a' + 10;
            } else {
                throw new ParseError("invalid char in color str");
            }

            value = value * 16;
            value += v;
        }

        return value;
    }

    private static Map<String, Color> colorMap = new HashMap<>();

    static {
        colorMap.put("black", Color.BLACK);
        colorMap.put("cyan", Color.CYAN);
        colorMap.put("blue", Color.BLUE);
        colorMap.put("darkgray", Color.DARK_GRAY);
        colorMap.put("gray", Color.GRAY);
        colorMap.put("green", Color.GREEN);
        colorMap.put("yellow", Color.YELLOW);
        colorMap.put("lightgray", Color.LIGHT_GRAY);
        colorMap.put("magenta", Color.MAGENTA);
        colorMap.put("orange", Color.ORANGE);
        colorMap.put("pink", Color.PINK);
        colorMap.put("red", Color.RED);
        colorMap.put("white", Color.WHITE);
    }

    private static Color parseName(String str) throws ParseError {
        if (colorMap.containsKey(str)) {
            return colorMap.get(str);
        } else {
            throw new ParseError("no color with specified name can be found");
        }
    }
}
