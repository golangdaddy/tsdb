package metrick

import (
	"testing"
)

func TestComputeSegmentQDAY(t *testing.T) {

	x := computeSegment(0, CONST_MEMBERS_DAY, CONST_SECTIONS_QDAY)
	if x != 0 {
		t.Errorf("INVALID SEGMENT: %v", x)
	}
	x = computeSegment(3, CONST_MEMBERS_DAY, CONST_SECTIONS_QDAY)
	if x != 0 {
		t.Errorf("INVALID SEGMENT: %v", x)
	}
	x = computeSegment(6, CONST_MEMBERS_DAY, CONST_SECTIONS_QDAY)
	if x != 1 {
		t.Errorf("INVALID SEGMENT: %v", x)
	}
	x = computeSegment(9, CONST_MEMBERS_DAY, CONST_SECTIONS_QDAY)
	if x != 1 {
		t.Errorf("INVALID SEGMENT: %v", x)
	}
	x = computeSegment(12, CONST_MEMBERS_DAY, CONST_SECTIONS_QDAY)
	if x != 2 {
		t.Errorf("INVALID SEGMENT: %v", x)
	}

}

func TestComputeSegmentQHOUR(t *testing.T) {

	x := computeSegment(0, CONST_MEMBERS_HOUR, CONST_SECTIONS_QHOUR)
	if x != 0 {
		t.Errorf("INVALID SEGMENT: %v", x)
	}
	x = computeSegment(14, CONST_MEMBERS_HOUR, CONST_SECTIONS_QHOUR)
	if x != 0 {
		t.Errorf("INVALID SEGMENT: %v", x)
	}
	x = computeSegment(15, CONST_MEMBERS_HOUR, CONST_SECTIONS_QHOUR)
	if x != 1 {
		t.Errorf("INVALID SEGMENT: %v", x)
	}
	x = computeSegment(29, CONST_MEMBERS_HOUR, CONST_SECTIONS_QHOUR)
	if x != 1 {
		t.Errorf("INVALID SEGMENT: %v", x)
	}
	x = computeSegment(30, CONST_MEMBERS_HOUR, CONST_SECTIONS_QHOUR)
	if x != 2 {
		t.Errorf("INVALID SEGMENT: %v", x)
	}
	x = computeSegment(44, CONST_MEMBERS_HOUR, CONST_SECTIONS_QHOUR)
	if x != 2 {
		t.Errorf("INVALID SEGMENT: %v", x)
	}
	x = computeSegment(45, CONST_MEMBERS_HOUR, CONST_SECTIONS_QHOUR)
	if x != 3 {
		t.Errorf("INVALID SEGMENT: %v", x)
	}
	x = computeSegment(59, CONST_MEMBERS_HOUR, CONST_SECTIONS_QHOUR)
	if x != 3 {
		t.Errorf("INVALID SEGMENT: %v", x)
	}

}

func TestComputeSegmentQMINUTE(t *testing.T) {

	x := computeSegment(0, CONST_MEMBERS_MINUTE, CONST_SECTIONS_QMINUTE)
	if x != 0 {
		t.Errorf("INVALID SEGMENT: %v", x)
	}
	x = computeSegment(14, CONST_MEMBERS_MINUTE, CONST_SECTIONS_QMINUTE)
	if x != 0 {
		t.Errorf("INVALID SEGMENT: %v", x)
	}
	x = computeSegment(15, CONST_MEMBERS_MINUTE, CONST_SECTIONS_QMINUTE)
	if x != 1 {
		t.Errorf("INVALID SEGMENT: %v", x)
	}
	x = computeSegment(29, CONST_MEMBERS_MINUTE, CONST_SECTIONS_QMINUTE)
	if x != 1 {
		t.Errorf("INVALID SEGMENT: %v", x)
	}
	x = computeSegment(30, CONST_MEMBERS_MINUTE, CONST_SECTIONS_QMINUTE)
	if x != 2 {
		t.Errorf("INVALID SEGMENT: %v", x)
	}
	x = computeSegment(44, CONST_MEMBERS_MINUTE, CONST_SECTIONS_QMINUTE)
	if x != 2 {
		t.Errorf("INVALID SEGMENT: %v", x)
	}
	x = computeSegment(45, CONST_MEMBERS_MINUTE, CONST_SECTIONS_QMINUTE)
	if x != 3 {
		t.Errorf("INVALID SEGMENT: %v", x)
	}
	x = computeSegment(59, CONST_MEMBERS_MINUTE, CONST_SECTIONS_QMINUTE)
	if x != 3 {
		t.Errorf("INVALID SEGMENT: %v", x)
	}

}
