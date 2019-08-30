package chrisww.watermark;

public class ExtSplitter {
    private String basename;
    private String extname;

    public String getBasename() {
        return basename;
    }

    public String getExtname() {
        return extname;
    }

    public ExtSplitter(String filename) {
        int dotIndex = filename.lastIndexOf('.');
        this.basename = filename.substring(0, dotIndex);
        this.extname = filename.substring(dotIndex + 1);
    }
}
