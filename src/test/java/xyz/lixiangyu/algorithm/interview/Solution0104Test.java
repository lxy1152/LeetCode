package xyz.lixiangyu.algorithm.interview;

import org.junit.Assert;
import org.junit.Test;
import xyz.lixiangyu.algorithm.AbstractTest;
import xyz.lixiangyu.algorithm.interview.string.Solution0104;

/**
 * {@link Solution0104}的测试类
 *
 * @author lixiangyu
 */
public class Solution0104Test extends AbstractTest {
    /**
     * {@inheritDoc}
     */
    @Override
    public void beforeTestMethod() {
        super.beforeTestMethod();
        solution = new Solution0104();
    }

    @Test
    public void testSolutionCaseOne() {
        Assert.assertEquals(true, solution.solve("tactcoa"));
    }

    @Test
    public void testSolutionCaseTwo() {
        Assert.assertEquals(false, solution.solve("tattcoa"));
    }
}
