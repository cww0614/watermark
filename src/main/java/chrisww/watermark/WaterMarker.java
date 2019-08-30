package chrisww.watermark;

import javax.imageio.ImageIO;
import java.awt.*;
import java.awt.geom.AffineTransform;
import java.awt.geom.Rectangle2D;
import java.awt.image.BufferedImage;
import java.io.File;
import java.io.IOException;

public class WaterMarker {
    private String text;

    private int horizontalSpacing;

    private int verticalSpacing;

    private float transparency;

    private int fontSize;

    private String fontName;

    private Color color;

    public void mark(File inputFile, File outputFile) throws IOException {
        BufferedImage sourceImage = ImageIO.read(inputFile);
        Graphics2D g2d = (Graphics2D) sourceImage.getGraphics();

        try {
            AlphaComposite alphaChannel = AlphaComposite.getInstance(AlphaComposite.SRC_OVER, transparency);
            g2d.setComposite(alphaChannel);
            g2d.setColor(color);
            g2d.setFont(new Font(fontName, Font.BOLD, fontSize));

            FontMetrics fontMetrics = g2d.getFontMetrics();
            Rectangle2D rect = fontMetrics.getStringBounds(text, g2d);

            int textBoxWidth = (int) (rect.getWidth() / Math.sqrt(2)) + horizontalSpacing * 2;
            int textBoxHeight = (int) (rect.getHeight()) + verticalSpacing * 2;

            for (int xOffset = -textBoxWidth; xOffset < sourceImage.getWidth() + textBoxWidth; xOffset += textBoxWidth) {
                for (int yOffset = -textBoxHeight; yOffset < sourceImage.getHeight() + textBoxHeight; yOffset += textBoxHeight) {
                    int centerX = (int) (xOffset + rect.getWidth() / 2);
                    int centerY = (int) (yOffset + rect.getHeight() / 2);

                    AffineTransform orig = g2d.getTransform();
                    g2d.rotate(Math.PI / 4, centerX, centerY);
                    g2d.drawString(text, xOffset, yOffset);
                    g2d.setTransform(orig);
                }
            }

            ImageIO.write(sourceImage, new ExtSplitter(outputFile.getName()).getExtname(), outputFile);
        } finally {
            g2d.dispose();
        }
    }

    public void setText(String text) {
        this.text = text;
    }

    public void setHorizontalSpacing(int horizontalSpacing) {
        this.horizontalSpacing = horizontalSpacing;
    }

    public void setVerticalSpacing(int verticalSpacing) {
        this.verticalSpacing = verticalSpacing;
    }


    public void setTransparency(float transparency) {
        this.transparency = transparency;
    }

    public void setFontSize(int fontSize) {
        this.fontSize = fontSize;
    }

    public void setFontName(String fontName) {
        this.fontName = fontName;
    }

    public void setColor(Color color) {
        this.color = color;
    }

    public String getText() {
        return text;
    }

    public int getHorizontalSpacing() {
        return horizontalSpacing;
    }

    public int getVerticalSpacing() {
        return verticalSpacing;
    }

    public float getTransparency() {
        return transparency;
    }

    public int getFontSize() {
        return fontSize;
    }

    public String getFontName() {
        return fontName;
    }

    public Color getColor() {
        return color;
    }

}
