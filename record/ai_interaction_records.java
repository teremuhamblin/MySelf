package myself.records;

import java.time.LocalDateTime;

public record AiInteractionRecord(
        String sender,           // "user" ou "assistant"
        String message,
        String intent,           // ex: "question", "command", "context"
        LocalDateTime timestamp
) {}
