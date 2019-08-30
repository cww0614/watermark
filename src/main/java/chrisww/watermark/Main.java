package chrisww.watermark;

import chrisww.watermark.colorparser.ColorParser;
import chrisww.watermark.colorparser.ParseError;
import picocli.CommandLine;
import picocli.CommandLine.Option;
import picocli.CommandLine.Parameters;

import java.io.File;
import java.util.concurrent.Callable;

@CommandLine.Command(name = "watermarker", mixinStandardHelpOptions = true, version = "1.0")
public class Main implements Callable<Integer> {
    @Option(names = {"-s", "--scale"},
            description = "Scale watermarks (default: ${DEFAULT-VALUE})")
    private double scale = 1.0;

    @Option(names = {"-o", "--horizontal-spacing"},
            description = "Horizontal spacing between watermarks (default: ${DEFAULT-VALUE})")
    private int horizontalSpacing = 40;

    @Option(names = {"-e", "--vertical-spacing"},
            description = "Vertical spacing between watermarks (default: ${DEFAULT-VALUE})")
    private int verticalSpacing = 40;

    @Option(names = {"-t", "--transparency"},
            description = "Transparency of watermark (default: ${DEFAULT-VALUE})")
    private float transparency = 0.05f;

    @Option(names = {"-f", "--font"},
            description = "Font for watermark text (default: ${DEFAULT-VALUE})")
    private String fontName = "Arial";

    @Option(names = {"-S", "--font-size"},
            description = "Font size for watermark text (default: ${DEFAULT-VALUE})")
    private int fontSize = 64;

    @Option(names = {"-c", "--color"},
            description = "Color for watermark text, name of #rrggbb (default: ${DEFAULT-VALUE})")
    private String color = "blue";

    @Parameters(index = "0", paramLabel = "TEXT",
            description = "Watermark text")
    private String text;

    @Parameters(index = "1..*", arity = "1..*", paramLabel = "FILE",
            description = "File(s) to process.")
    private File[] inputFiles;

    private WaterMarker createWaterMarker() throws ParseError {
        WaterMarker marker = new WaterMarker();
        marker.setText(text);

        marker.setHorizontalSpacing(horizontalSpacing);
        marker.setVerticalSpacing(verticalSpacing);
        marker.setTransparency(transparency);
        marker.setFontName(fontName);
        marker.setFontSize(fontSize);
        marker.setColor(ColorParser.parse(color));

        marker.setFontSize((int) (marker.getFontSize() * scale));
        marker.setHorizontalSpacing((int) (marker.getHorizontalSpacing() * scale));
        marker.setVerticalSpacing((int) (marker.getVerticalSpacing() * scale));

        return marker;
    }

    @Override
    public Integer call() {
        try {
            WaterMarker marker = createWaterMarker();

            for (File inputFile : inputFiles) {
                if (!inputFile.exists()) {
                    System.err.println(String.format("File %s doesn't exists!", inputFile.getAbsolutePath()));
                    return 1;
                }

                ExtSplitter extSplitter = new ExtSplitter(inputFile.getName());

                File outputFile = new File(inputFile.getParentFile(),
                        String.format("%s.watermarked.%s", extSplitter.getBasename(), extSplitter.getExtname()));

                marker.mark(inputFile, outputFile);
            }

            return 0;
        } catch (Exception e) {
            e.printStackTrace();
            return 1;
        }
    }

    public static void main(String[] args) {
        int exitCode = new CommandLine(new Main()).execute(args);
        System.exit(exitCode);
    }
}
