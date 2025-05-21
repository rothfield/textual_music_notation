package newparser

import (
    "fmt"
    "strings"
)

type StringFormatter struct {
    Builder strings.Builder
}

func (f *StringFormatter) WriteLine(indent, line string) {
    f.Builder.WriteString(indent + line + "\n")
}

func RenderLetterLine(line *LetterLine, formatter *StringFormatter, indent string) {
    formatter.WriteLine(indent, "LetterLine")
    for _, el := range line.Elements {
        if el.IsBeat {
            beatStr := ""
            for _, sub := range el.SubElements {
                beatStr += sub.Token.Value
                if sub.SyllableText != "" {
                    beatStr += fmt.Sprintf(" [%s]", sub.SyllableText)
                }
            }
            formatter.WriteLine(indent+"  ", "- Beat: "+beatStr)
        } else {
            itemStr := el.Token.Value
            if el.SyllableText != "" {
                itemStr += fmt.Sprintf(" [%s]", el.SyllableText)
            }
            formatter.WriteLine(indent+"  ", fmt.Sprintf("- %s: %s", el.Token.Type, itemStr))
        }
    }
}

