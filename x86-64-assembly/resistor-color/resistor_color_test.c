// Version: 1.0.0

#include "vendor/unity.h"

#define ARRAY_SIZE(x) (sizeof(x) / sizeof((x)[0]))

extern int color_code(const char *color);
extern const char **colors(void);
extern int string_compare(const char *a, const char *b);

void setUp(void) {
}

void tearDown(void) {
}

void test_string_compare(void) {
    TEST_ASSERT_EQUAL_INT(1, string_compare("a", "a"));
    TEST_ASSERT_EQUAL_INT(0, string_compare("a", "b"));
    TEST_ASSERT_EQUAL_INT(0, string_compare("a1", "a2"));
    TEST_ASSERT_EQUAL_INT(1, string_compare("abcd", "abcd"));
    TEST_ASSERT_EQUAL_INT(0, string_compare("abcd", "ab"));
}

void test_black(void) {
    TEST_ASSERT_EQUAL_INT(0, color_code("black"));
}

void test_white(void) {
    TEST_ASSERT_EQUAL_INT(9, color_code("white"));
}

void test_orange(void) {
    TEST_ASSERT_EQUAL_INT(3, color_code("orange"));
}

void test_colors(void) {
    const char **color_array = colors();
    const char *expected[] = {"black", "brown", "red", "orange", "yellow", "green", "blue", "violet", "grey", "white"};
    int size;

    for (size = 0; color_array[size]; size++) {
    }
    TEST_ASSERT_EQUAL_INT(ARRAY_SIZE(expected), size);
    TEST_ASSERT_EQUAL_STRING_ARRAY(expected, color_array, size);
}

int main(void) {
    UNITY_BEGIN();
    RUN_TEST(test_string_compare);
    RUN_TEST(test_black);
    RUN_TEST(test_white);
    RUN_TEST(test_orange);
    RUN_TEST(test_colors);
    return UNITY_END();
}
