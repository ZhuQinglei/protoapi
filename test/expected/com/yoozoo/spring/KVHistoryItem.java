// Code generated by protoapi; DO NOT EDIT.

package com.yoozoo.spring;

import com.fasterxml.jackson.annotation.JsonCreator;
import com.fasterxml.jackson.annotation.JsonProperty;

import java.util.List;

public class KVHistoryItem {
    private final String updated_value;
    private final String updated_date;
    private final String updated_by;
    private final long revision;

    @JsonCreator
    public KVHistoryItem(@JsonProperty("updated_value") String updated_value, @JsonProperty("updated_date") String updated_date, @JsonProperty("updated_by") String updated_by, @JsonProperty("revision") long revision) {
        this.updated_value = updated_value;
        this.updated_date = updated_date;
        this.updated_by = updated_by;
        this.revision = revision;
    }

    public String getUpdated_value() {
        return updated_value;
    }
    public String getUpdated_date() {
        return updated_date;
    }
    public String getUpdated_by() {
        return updated_by;
    }
    public long getRevision() {
        return revision;
    }
    
}