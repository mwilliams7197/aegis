ALTER TABLE AssignmentRules ADD COLUMN ExcludeVulnTitleRegex VARCHAR(300) NULL AFTER VulnTitleRegex;
ALTER TABLE AssignmentRules ADD COLUMN OSRegex VARCHAR(300) NULL AFTER HostnameRegex;
ALTER TABLE AssignmentRules ADD COLUMN PortCSV VARCHAR(300) NULL AFTER TagKeyRegex;
ALTER TABLE AssignmentRules ADD COLUMN ExcludePortCSV VARCHAR(300) NULL AFTER PortCSV;