package myself.records;

public record IdentityRecord(
        String fullName,
        String role,
        String location,
        String background,
        String[] skills,
        String[] interests
) {}
